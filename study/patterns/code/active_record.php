<?php

class ActiveRecord
{
    protected PDO $db;
    public string $id;
    public string $bar;

    public function __construct(PDO $db)
    {
        $this->db = $db;
    }

    public function do_something(): void
    {
        $this->bar .= uniqid();
    }

    public function save(): void
    {
        if ($this->id) {
            $sql = "UPDATE foo SET bar = :bar WHERE id = :id";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $this->bar);
            $statement->bindParam("id", $this->id);
            $statement->execute();
        }
        else {
            $sql = "INSERT INTO foo (bar) VALUES (:bar)";
            $statement = $this->db->prepare($sql);
            $statement->bindParam("bar", $this->bar);
            $statement->execute();
            $this->id = $this->db->lastInsertId();
        }
    }
}

$db = new PDO("sqlite:foo.db");
//Insert
$foo = new ActiveRecord($db);
$foo->bar = 'baz';
$foo->save();