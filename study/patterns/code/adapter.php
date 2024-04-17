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

interface EBook
{
    public function unlock();
    public function pressNext();
    public function getPage(): array;
}

class EBookAdapter implements Book
{
    protected EBook $eBook;
    public function __construct(EBook $eBook)
    {
        $this->eBook = $eBook;
    }
    public function open(): void
    {
        $this->eBook->unlock();
    }
    public function turnPage(): void
    {
        $this->eBook->pressNext();
    }
    public function getPage(): int
    {
        return $this->eBook->getPage()[0];
    }
}

class Kindle implements EBook
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

$kindle = new Kindle();
$book = new EBookAdapter($kindle);
$book->open();
$book->turnPage();
