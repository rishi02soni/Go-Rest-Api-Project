# 🚀 Go Task Manager API

POST /tasks
```

### 🔹 Get Task by ID
```
GET /tasks/{id}
```

### 🔹 Update Task
```
PUT /tasks/{id}
```

### 🔹 Delete Task
```
DELETE /tasks/{id}
```

---

## 📥 Sample Request Body

```json
{
  "id": "1",
  "title": "Learn Go",
  "done": false
}
```

---

## 🧪 Testing the API

You can test APIs using:
- Postman
- Thunder Client (VS Code Extension)
- curl (CLI)

Example using curl:
```
curl -X GET http://localhost:8080/tasks
```

---

## ⚠️ Note

- This project uses **in-memory storage**, so data will reset when server restarts.
- It is meant for learning purposes.

---

## 🚀 Future Improvements

- Add Database (MongoDB / MySQL)
- Add Authentication (JWT)
- Add Validation & Middleware
- Dockerize the application
- Deploy on cloud (AWS / Render / Railway)

---

## 👨‍💻 Author

Rishi Soni

---

## ⭐ Support

If you like this project, give it a ⭐ on GitHub and share it with others!


