<?php

use common\components\helpers\ArrayHelper;
use common\models\order\LocalOrderProvider;
use common\models\order\PSCBOrderProvider;
use exclusive\pscb\PSCBModule;
use exclusive\pscb\source\mvc\models\payment\FdReceipt;
use exclusive\pscb\source\mvc\models\payment\Item;
use exclusive\pscb\source\mvc\models\payment\Payment;
use exclusive\pscb\source\PSCBModuleConfig;
use yii\base\Module;
use yii\web\Application;
use yii\web\UrlManager;

$scheme = 'http';

return [
    'class' => PSCBModule::class,
    'on ' . Module::EVENT_BEFORE_ACTION => function () use ($scheme) {
        /**
         * @var UrlManager $urlManagerApi
         * @var UrlManager $urlManagerFrontend
         */
        $urlManagerApi = Yii::$app->urlManager;
        $urlManagerFrontend = Yii::$app->urlManagerFrontend;
        $config = [
            'kibana' => [
                'index' => 'order_api_test'
            ],
            'api' => [
                'rootUrl' => 'https://oosdemo.pscb.ru',
                'marketPlaceId' => '144611094',
                'merchantKey' => '111111',
                'companyEmail' => 'ooleg.chern@gmail.com'
            ],
            'orderProvider' => [
                'pscb' => PSCBOrderProvider::class,
                'orderApi' => LocalOrderProvider::class
            ],
            'makePaymentOptions' => [
                'recurrentable' => false,
                'hold' => true,
                'debug' => true,
                'taxSystem' => FdReceipt::TAX_SYSTEM,
                'itemTax' => Item::TAX,
                'paymentType' => Payment::FULL_PREPAYMENT
            ]
        ];
        if (Yii::$app instanceof Application) {
            $config = ArrayHelper::merge(
                $config,
                [
                    'startPaymentUrl' => $urlManagerApi->createAbsoluteUrl([
                        '/pscb/pscb-web/payment-request',
                        'order' => '~order~'
                    ], $scheme),
                    'pscbRedirectUrls' => [
                        'success' => $urlManagerApi
                            ->createAbsoluteUrl('/pscb/pscb-web/payment-response-success', $scheme),
                        'fail' =>
                            $urlManagerApi->createAbsoluteUrl('/pscb/pscb-web/payment-response-fail', $scheme),
                    ],
                    'tyresRedirectUrls' => [
                        'paymentResponseSuccessUrl' => $urlManagerFrontend
                            ->createAbsoluteUrl('/cart_success', $scheme),
                        'paymentResponseFailUrl' => $urlManagerFrontend
                            ->createAbsoluteUrl('/cart_order', $scheme),
                        'cartUrl' => $urlManagerFrontend->createAbsoluteUrl('/cart_order', $scheme),
                    ]
                ]
            );
        }
        PSCBModuleConfig::create($config);
    }
];