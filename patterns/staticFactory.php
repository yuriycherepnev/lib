<?php

use InvalidArgumentException;

/**
 * Note1: Remember, static means global state which is evil because it can't be mocked for tests
 * Note2: Cannot be subclassed or mock-upped or have multiple different instances.
 */
final class StaticFactory
{
    public static function factory(string $type): Formatter
    {
        return match ($type) {
            'number' => new FormatNumber(),
            'string' => new FormatString(),
            default => throw new InvalidArgumentException('Unknown format given'),
        };
    }
}

interface Formatter
{
    public function format(string $input): string;
}

class FormatString implements Formatter
{
    public function format(string $input): string
    {
        return $input;
    }
}

class FormatNumber implements Formatter
{
    public function format(string $input): string
    {
        return number_format((int) $input);
    }
}


StaticFactory::factory('number');
StaticFactory::factory('string');
StaticFactory::factory('object');
