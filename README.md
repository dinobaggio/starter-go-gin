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