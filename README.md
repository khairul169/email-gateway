# Email Gateway API

This is a simple email-sending gateway built with Go and go-mail. It allows sending emails with attachments via multiple SMTP providers, secured with API keys.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/khairul169/email-gateway.git
   cd email-gateway
   ```
2. Install dependencies:
   ```sh
   go mod download
   ```
3. Start the server
   ```sh
   go run main.go
   ```

## API Endpoint

### `POST /send-email`

#### Headers:

- `X-API-Key`: Your API key

#### Form Data Parameters:

| Parameter  | Type   | Required | Description                 |
| ---------- | ------ | -------- | --------------------------- |
| email      | string | Yes      | Recipient's email address   |
| title      | string | Yes      | Email subject               |
| content    | string | Yes      | Email body (HTML supported) |
| attachment | file   | No       | File attachment             |

## Sample API Calls

### Using PHP (cURL)

```php
$ch = curl_init('https://example.com/send-email');
$postData = [
    'email' => 'recipient@example.com',
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
const formData = new FormData();
formData.append("email", "recipient@example.com");
formData.append("title", "Test Email");
formData.append("content", "<h1>Hello</h1><p>This is a test email.</p>");

const attachment = fs.readFileSync("./path/to/file.pdf");
formData.append("attachment", new Blob([attachment]), "file.pdf"); // Optional file

fetch("https://example.com/send-email", {
  method: "POST",
  headers: { "X-API-Key": "your-api-key-1" },
  body: formData,
})
  .then((response) => response.json())
  .then((data) => console.log(data))
  .catch((error) => console.error("Error:", error));
```

## License

This project is licensed under the MIT License.
