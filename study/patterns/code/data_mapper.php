<?php

class User
{
    public static function fromState(array $state): User
    {
        return new self(
            $state['username'],
            $state['email']
        );
    }
    public function __construct(private string $username, private string $email)
    {
    }
    public function getUsername(): string
    {
        return $this->username;
    }
    public function getEmail(): string
    {
        return $this->email;
    }
}

class UserMapper
{
    public function findById(int $id): ?User
    {
        /* sql query */
        if (isset($this->data[$id])) {
            return $this->mapRowToUser($this->data[$id]);
        }
        return null;
    }
    private function mapRowToUser(array $row): User
    {
        return User::fromState($row);
    }
}

$mapper = new UserMapper();
$user = $mapper->findById(1);
echo $user->getUsername();
echo $user->getEmail();