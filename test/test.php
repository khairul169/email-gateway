<?php

$ch = curl_init('http://localhost:8080');
$postData = [
    'email' => 'vobisem902@bankrau.com',
    'name' => 'Recipient Name',
    'title' => 'Test Email',
    'content' => '<h1>Hello</h1><p>This is a test email.</p>',
    'attachment' => new CURLFile('./attachment.txt') // Optional file
];

curl_setopt($ch, CURLOPT_HTTPHEADER, ['X-API-Key: test']);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);
if ($response === false) {
    $error = curl_error($ch);
    curl_close($ch);
    die("cURL Error: " . $error);
}

$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

// Handle HTTP response codes
if ($httpCode >= 400) {
    die("API Error: HTTP $httpCode - Response: $response");
}

echo $response;
