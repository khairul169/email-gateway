import fs from "fs";

async function send() {
  const formData = new FormData();
  formData.append("email", "vobisem902@bankrau.com");
  formData.append("name", "Recipient Name");
  formData.append("title", "Test Email");
  formData.append("content", "<h1>Hello</h1><p>This is a test email.</p>");

  const attachment = fs.readFileSync("./attachment.txt");
  formData.append("attachment", new Blob([attachment]), "attachment.txt"); // Optional file

  try {
    const res = await fetch("http://localhost:8080", {
      method: "POST",
      headers: { "x-api-key": "test" },
      body: formData,
    });
    const data = await res.json();
    console.log(data);
  } catch (err) {
    console.error(err);
  }
}

send();
