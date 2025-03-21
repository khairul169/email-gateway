<?php
require __DIR__ . '/vendor/autoload.php';
require __DIR__ . '/config.php';

use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\Exception;

// Only allow POST requests
if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    http_response_code(405);
    echo json_encode(['error' => 'Method not allowed']);
    exit;
}

// Get API Key
$headers = array_change_key_case(getallheaders(), CASE_LOWER);
$apiKey = $headers['x-api-key'] ?? $headers['authorization'] ?? '';

if (!isset($apiKeys[$apiKey])) {
    http_response_code(401);
    echo json_encode(['error' => 'Invalid API key']);
    exit;
}

$smtpConfig = $apiKeys[$apiKey];

// Get Form Data
$email = $_POST['email'] ?? '';
$name = $_POST['name'] ?? '';
$title = $_POST['title'] ?? '';
$content = $_POST['content'] ?? '';

if (!$email || !$title || !$content) {
    http_response_code(400);
    echo json_encode(['error' => 'Missing required fields']);
    exit;
}

$mail = new PHPMailer(true);
try {
    $mail->isSMTP();
    $mail->Host = $smtpConfig['host'];
    $mail->SMTPAuth = true;
    $mail->Username = $smtpConfig['username'];
    $mail->Password = $smtpConfig['password'];
    $mail->SMTPSecure = PHPMailer::ENCRYPTION_STARTTLS;
    $mail->Port = $smtpConfig['port'] ?? 587;

    $mail->setFrom($smtpConfig['username'], $smtpConfig['name'] ?? '');
    $mail->addAddress($email, $name);
    $mail->Subject = $title;
    $mail->isHTML(true);
    $mail->Body = $content;

    if (!empty($_FILES['attachment']['tmp_name'])) {
        $mail->addAttachment($_FILES['attachment']['tmp_name'], $_FILES['attachment']['name']);
    }

    $mail->send();
    echo json_encode(['success' => 'Email sent successfully']);
} catch (Exception $e) {
    http_response_code(500);
    echo json_encode(['error' => "Email could not be sent. Mailer Error: {$mail->ErrorInfo}"]);
}
