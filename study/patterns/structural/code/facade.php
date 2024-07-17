<?php

class Facade
{
    public function __construct(private Bios $bios, private OperatingSystem $os)
    {
    }

    public function turnOn(): void
    {
        $this->bios->execute();
        $this->bios->waitForKeyPress();
        $this->bios->launch($this->os);
    }

    public function turnOff(): void
    {
        $this->os->halt();
        $this->bios->powerDown();
    }
}

interface OperatingSystem
{
    public function halt();
}

interface Bios
{
    public function execute();
    public function waitForKeyPress();
    public function launch(OperatingSystem $os);
    public function powerDown();
}