<?php

interface Cake
{
    public function ingredients(): string;
    public function price(): float;
}

abstract class CakeDecorator implements Cake
{
    protected Cake $cake;
    public function __construct(Cake $cake)
    {
        $this->cake = $cake;
    }
}

class SimpleCake implements Cake
{
    public function ingredients(): string
    {
        return 'Простой торт';
    }
    public function price(): float
    {
        return 12.5;
    }
}

class WithWhippedCream extends CakeDecorator
{
    public function ingredients(): string
    {
        return $this->cake->ingredients() . ' со взбитыми сливками';
    }
    public function price(): float
    {
        return $this->cake->price() + 2.5;
    }
}

class WithSprinkles extends CakeDecorator
{
    public function ingredients(): string
    {
        return $this->cake->ingredients() . ' с обсыпкой';
    }
    public function price(): float
    {
        return $this->cake->price() + 1.25;
    }
}

$cake = new SimpleCake();
$cake = new WithWhippedCream($cake);
$cake = new WithSprinkles($cake);
echo $cake->ingredients(); // Простой торт со взбитыми сливками с обсыпкой
echo $cake->price(); // 16.25
// или можно цепочкой
$cake = new WithSprinkles(new WithWhippedCream(new SimpleCake()));
echo $cake->ingredients(); // Простой торт со взбитыми сливками с обсыпкой
echo $cake->price(); // 16.25