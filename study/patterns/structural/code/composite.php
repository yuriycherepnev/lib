<?php

interface Renderer
{
    public function render(): string;
}

class Form implements Renderer
{
    private array $elements;
    public function render(): string
    {
        $formCode = '<form>';
        foreach ($this->elements as $element) {
            $formCode .= $element->render();
        }
        return $formCode . '</form>';
    }
    public function addElement(Renderer $element): void
    {
        $this->elements[] = $element;
    }
}

class InputElement implements Renderer
{
    public function render(): string
    {
        return '<input type="text" />';
    }
}

class TextElement implements Renderer
{
    private string $text;
    public function __construct(string $text)
    {
        $this->text = $text;
    }
    public function render(): string
    {
        return $this->text;
    }
}

$form = new Form();
$form->addElement(new TextElement('Email:'));
$form->addElement(new InputElement());
$embed = new Form();
$embed->addElement(new TextElement('Password:'));
$embed->addElement(new InputElement());
$form->addElement($embed);