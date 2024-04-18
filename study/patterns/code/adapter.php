<?php

interface Book
{
    public function turnPage();
    public function open();
    public function getPage(): int;
}

class PaperBook implements Book
{
    private int $page;
    public function open(): void
    {
        $this->page = 1;
    }
    public function turnPage(): void
    {
        $this->page++;
    }
    public function getPage(): int
    {
        return $this->page;
    }
}

interface Scroll
{
    public function unlock();
    public function pressNext();
    public function getPage(): array;
}

class BookAdapter implements Book
{
    protected Scroll $scroll;
    public function __construct(Scroll $scroll)
    {
        $this->scroll = $scroll;
    }
    public function open(): void
    {
        $this->scroll->unlock();
    }
    public function turnPage(): void
    {
        $this->scroll->pressNext();
    }
    public function getPage(): int
    {
        return $this->scroll->getPage()[0];
    }
}

class OldScroll implements Scroll
{
    private int $page = 1;
    private int $totalPages = 100;
    public function pressNext(): void
    {
        $this->page++;
    }
    public function unlock()
    {
    }
    public function getPage(): array
    {
        return [$this->page, $this->totalPages];
    }
}

$book = new PaperBook();
$book->open();
$book->turnPage();

$oldScroll = new OldScroll();
$book = new BookAdapter($oldScroll);
$book->open();
$book->turnPage();
