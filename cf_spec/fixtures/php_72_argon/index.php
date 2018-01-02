<h1>Password encrypt 'test' using argon2i</h1>
<?php
$password = 'test';
$hash = password_hash($password, PASSWORD_ARGON2I);
var_dump($hash);
?>
