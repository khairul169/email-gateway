# PHP Email Gateway API

This is a simple email-sending gateway built with PHP and PHPMailer. It allows sending emails with attachments via multiple SMTP providers, secured with API keys.

## Features

- Secure email sending using API keys
- Support for multiple SMTP providers
- Uses PHPMailer for email handling

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/email-gateway.git
   cd email-gateway
   ```
2. Install dependencies using Composer:
   ```sh
   composer install
   ```
3. Deploy the PHP script to your server.

## API Endpoint

### `POST /index.php`

#### Headers:

- `X-API-Key`: Your API key

#### Form Data Parameters:

| Parameter  | Type   | Required | Description                 |
| ---------- | ------ | -------- | --------------------------- |
| email      | string | Yes      | Recipient's email address   |
| name       | string | No       | Recipient's name            |
| title      | string | Yes      | Email subject               |
| content    | string | Yes      | Email body (HTML supported) |
| attachment | file   | No       | File attachment             |

## Sample API Calls

### Using PHP (cURL)

```php
$ch = curl_init('https://example.com/index.php');
$postData = [
    'email' => 'recipient@example.com',
    'name' => 'Recipient Name',
    'title' => 'Test Email',
    'content' => '<h1>Hello</h1><p>This is a test email.</p>',
    'attachment' => new CURLFile('/path/to/file.pdf') // Optional file
];

curl_setopt($ch, CURLOPT_HTTPHEADER, ['X-API-Key: your-api-key-1']);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

$response = curl_exec($ch);
curl_close($ch);

echo $response;
```

### Using Node.js (Fetch + FormData)

```javascript
const fs = require("fs");
const FormData = require("form-data");
const fetch = require("node-fetch");

const formData = new FormData();
formData.append("email", "recipient@example.com");
formData.append("name", "Recipient Name");
formData.append("title", "Test Email");
formData.append("content", "<h1>Hello</h1><p>This is a test email.</p>");

const attachment = fs.readFileSync("./path/to/file.pdf");
formData.append("attachment", new Blob([attachment]), "file.pdf"); // Optional file

fetch("https://example.com/index.php", {
  method: "POST",
  headers: {
    "X-API-Key": "your-api-key-1",
    ...formData.getHeaders(),
  },
  body: formData,
})
  .then((response) => response.json())
  .then((data) => console.log(data))
  .catch((error) => console.error("Error:", error));
```

## License

This project is licensed under the MIT License.
