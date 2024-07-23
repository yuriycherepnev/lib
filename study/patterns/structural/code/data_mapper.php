<?php

class Data
{
    public string $id;
    public string $bar;
}

class DataMapper
{
    protected PDO $db;

    public function __construct(PDO $db)
    {
        $this->db = $db;
    }

    public function saveData(Data &$data): void
    {
        if ($data->id) {
            $sql = "UPDATE foo SET bar = :bar WHERE id = :id";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $data->bar);
            $statement->bindParam("id", $data->id);
            $statement->execute();
        } else {
            $sql = "INSERT INTO foo (bar) VALUES (:bar)";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $data->bar);
            $statement->execute();
            $data->id = $this->db->lastInsertId();
        }
    }
}
$db = new PDO("sqlite:foo.db");
//Insert
$foo = new Data();
$foo->bar = 'baz';

$mapper = new DataMapper($db);
$mapper->saveData($foo);