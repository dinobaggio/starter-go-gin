## **Starter Golang with GIN**

- [x] Gin
- [x] Database SQL
- [x] Goose
- [x] Dockerize
- [x] Migrations
- [x] Cobra
- [x] Unit test
- [x] Login & Registration

## **How To Run**

## **Run Dev**

```bash
make dev
```

this command will be create docker instance main app and mysql database

## **Endpoint**

```bash
GET    /api/v1/healthcheck
POST   /api/v1/auth/login
POST   /api/v1/auth/register
GET    /api/v1/users
```

### **`GET` /api/v1/healthcheck**

**response :**
```json
"It's work"
```

### **`POST` /api/v1/auth/login**

**payload :**
```json
{
    "email": "admin@admin.com",
    "password": "password123",
}
```

**response :**
```json
{
    "message": "success",
    "data": {
        "token": "string",
        "user": {
            "id": "number",
            "uuid": "string",
            "name": "string",
            "email": "string",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
        }
    },
}
```

### **`POST` /api/v1/auth/register**

**payload :**
```json
{
    "name": "admin",
    "email": "admin@admin.com",
    "password": "password123",
    "confirm_password": "password123",
}
```

**response :**
```json
{
    "message": "success",
    "data": {
        "user": {
            "id": "number",
            "uuid": "string",
            "name": "string",
            "email": "string",
            "created_at": "string",
            "updated_at": "string",
            "deleted_at": "string",
        }
    },
}
```

### **`GET` /api/v1/auth/users**

**response :**
```json
{
    "message": "success",
    "data": [{
        "id": "number",
        "name": "string",
        "email": "string",
        "created_at": "string",
        "updated_at": "string",
    }],
}
```

### **Credentials**

email: `admin@admin.com`

password: `password123`


## **Migration**

## **How To Create Migration**

```bash
make create-migration FILENAME=create_table_users
```

output file will be on `database/migrations`

## **How To Run Migration**

```bash
go run . migrate
```

## **How To Down Migration**

```bash
go run . migrate:down
```