<?php

class Data
{
    public string $id;
    public string $bar;

    public function do_something(): void
    {
        $this->bar .= uniqid();
    }
}

class DataMapper
{
    protected PDO $db;

    public function __construct(PDO $db)
    {
        $this->db = $db;
    }

    public function saveFoo(Data &$foo): void
    {
        if ($foo->id) {
            $sql = "UPDATE foo SET bar = :bar WHERE id = :id";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $foo->bar);
            $statement->bindParam("id", $foo->id);
            $statement->execute();
        }
        else {
            $sql = "INSERT INTO foo (bar) VALUES (:bar)";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $foo->bar);
            $statement->execute();
            $foo->id = $this->db->lastInsertId();
        }
    }
}
$db = new PDO("sqlite:foo.db");
//Insert
$foo = new Data();
$foo->bar = 'baz';

$mapper = new DataMapper($db);
$mapper->saveFoo($foo);