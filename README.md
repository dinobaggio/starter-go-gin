# **Starter Golang with GIN**

- [x] Gin
- [x] Database SQL
- [x] Goose
- [x] Dockerize
- [x] Migrations
- [x] Cobra
- [x] Unit test
- [x] Login & Registration

# **How To Run**

## **Run Dev**

```bash
make dev
```

this command will be create docker instance main app and mysql database

## **Endpoint**

```bash
[GIN-debug] GET    /api/v1/healthcheck       --> starter-go-gin/routes.SetUp.func1 (3 handlers)
[GIN-debug] POST   /api/v1/auth/login        --> starter-go-gin/handlers.(*AuthHandler).Login-fm (3 handlers)
[GIN-debug] POST   /api/v1/auth/register     --> starter-go-gin/handlers.(*AuthHandler).Register-fm (3 handlers)
[GIN-debug] GET    /api/v1/users             --> starter-go-gin/handlers.(*UserHandler).List-fm (3 handlers)
```

### **Credentials**

email: `admin@admin.com`

password: `password123`


# **Migration**

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