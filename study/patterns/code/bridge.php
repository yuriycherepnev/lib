<?php

interface WebPage
{
    public function __construct(Theme $theme);
    public function getContent();
}

class About implements WebPage
{
    protected Theme $theme;
    public function __construct(Theme $theme)
    {
        $this->theme = $theme;
    }
    public function getContent(): string
    {
        return "About page in " . $this->theme->getColor();
    }
}

class Careers implements WebPage
{
    protected Theme $theme;
    public function __construct(Theme $theme)
    {
        $this->theme = $theme;
    }
    public function getContent(): string
    {
        return "Careers page in " . $this->theme->getColor();
    }
}

interface Theme
{
    public function getColor();
}

class DarkTheme implements Theme
{
    public function getColor(): string
    {
        return 'Dark Black';
    }
}

class LightTheme implements Theme
{
    public function getColor(): string
    {
        return 'Off white';
    }
}

class AquaTheme implements Theme
{
    public function getColor(): string
    {
        return 'Light blue';
    }
}

$darkTheme = new DarkTheme();

$about = new About($darkTheme);
$careers = new Careers($darkTheme);

echo $about->getContent();
echo $careers->getContent();