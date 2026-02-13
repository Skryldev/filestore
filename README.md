<div dir="rtl">

# 📦 پکیج FileStore – ماژول حرفه‌ای ذخیره‌سازی فایل با Golang و MinIO
### یک ماژول Production-Ready، مستقل از فریم‌ورک و قابل تعویض با هر Storage Provider برای مدیریت فایل‌ها در پروژه‌های Golang.
#### این ماژول بر اساس اصول Clean Architecture طراحی شده و امکان تعویض MinIO با هر Storage دیگری (S3, GCS, Azure) را فراهم می‌کند.
---
## ✨ ویژگی‌ها
#### 🏗 معماری حرفه‌ای
- کاملاً مستقل از Gin یا هر فریم‌ورک دیگر
- طراحی مبتنی بر Interface و Abstraction
- قابل جایگزینی با هر storage provider
- بدون وابستگی به HTTP layer
---
## 🆔 مدیریت حرفه‌ای شناسه‌ها
- استفاده از UUID برای نام فایل در Storage (جلوگیری از تداخل)
- استفاده از ULID برای شناسه public API (sortable و خواناتر)
---
## 📤 قابلیت‌های اصلی
- آپلود فایل (Streaming، بدون بارگذاری کامل در حافظه)
- دانلود امن فایل
- حذف فایل
- تولید Presigned URL با زمان انقضا
- لیست فایل‌ها با فیلتر prefix / suffix
- پشتیبانی از Versioning در صورت فعال بودن Bucket
---
## 🚀 عملکرد بالا
- استفاده کامل از context.Context
- پشتیبانی از concurrency بالا
- مصرف بهینه حافظه برای فایل‌های حجیم
- Retry mechanism
- Structured Logging
- قابل استفاده در محیط‌های High Throughput
---
## 🔐 امنیت
- پشتیبانی از Server-Side Encryption
- اعتبارسنجی اندازه و نوع فایل
- Presigned URL با مدت اعتبار مشخص
- جداسازی Domain Errors از Infra Errors
---
## 🛠 نصب و راه‌اندازی
```bash
go get github.com/Skryldev/filestore
go get github.com/minio/minio-go/v7
go get github.com/google/uuid
go get github.com/oklog/ulid/v2
```
---
## 1️⃣ ساخت Config

<div dir="ltr">

```go
cfg, err := filestore.LoadFromEnv()
if err != nil {
    log.Fatal(err)
}
```

<div dir="rtl">

##### یا ساخت دستی:

<div dir="ltr">

```go
cfg := &filestore.Config{
    Endpoint:  "localhost:9000",
    AccessKey: "minioadmin",
    SecretKey: "minioadmin",
    UseSSL:    false,
    Bucket:    "files",
}
```

<div dir="rtl">

---
# 🚀 استفاده کامل از ماژول
## 2️⃣ ساخت Logger

<div dir="ltr">

```go
logger := filestore.NewZapLogger()
```

<div dir="rtl">

## 3️⃣ ساخت Storage

<div dir="ltr">

```go
storage, err := minioadapter.New(cfg, logger)
if err != nil {
    log.Fatal(err)
}
```

<div dir="rtl">

## 4️⃣ آپلود فایل

<div dir="ltr">

```go
file, err := os.Open("test.jpg")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

stat, _ := file.Stat()

info, err := storage.Upload(
    context.Background(),
    file,
    stat.Size(),
    "test.jpg",
    filestore.UploadOptions{
        ContentType: "image/jpeg",
    },
)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Public ID:", info.ID)
```

<div dir="rtl">

## 5️⃣ دانلود فایل

<div dir="ltr">

```go
reader, meta, err := storage.Download(ctx, publicID)
if err != nil {
    log.Fatal(err)
}
defer reader.Close()

io.Copy(os.Stdout, reader)
```

<div dir="rtl">

## 6️⃣ حذف فایل

<div dir="ltr">

```go
err := storage.Delete(ctx, publicID)
if err != nil {
    log.Fatal(err)
}
```

<div dir="rtl">

## 7️⃣ تولید Presigned URL

<div dir="ltr">

```go
url, err := storage.PresignedURL(ctx, publicID, 15*time.Minute)
if err != nil {
    log.Fatal(err)
}

fmt.Println(url)
```

<div dir="rtl">

## 8️⃣ لیست فایل‌ها

<div dir="ltr">

```go
files, err := storage.List(ctx, filestore.ListOptions{
    Prefix: "images/",
})
if err != nil {
    log.Fatal(err)
}

for _, f := range files {
    fmt.Println(f.ID, f.Name)
}
```

<div dir="rtl">

---
## 🔌 اتصال به Gin (اختیاری)
##### این ماژول مستقل از Gin است، اما می‌توانید adaptor بسازید:

<div dir="ltr">

```go
func UploadHandler(storage filestore.Storage) gin.HandlerFunc {
return func(c *gin.Context) {
file, header, err := c.Request.FormFile("file")
if err != nil {
c.JSON(400, gin.H{"error": err.Error()})
return
    }
defer file.Close()

info, err := storage.Upload(
c.Request.Context(),
file,
header.Size,
header.Filename,
filestore.UploadOptions{
ContentType: header.Header.Get("Content-Type"),
    },
)

if err != nil {
c.JSON(500, gin.H{"error": err.Error()})
return
}

c.JSON(201, info)
    }
}
```

<div dir="rtl">

---
## ⚡ بهترین شیوه‌های استفاده در Production
* ✔ همیشه از Context با timeout استفاده کنید
* ✔ Bucket versioning را فعال کنید
* ✔ Server-Side Encryption را فعال کنید
* ✔ اندازه فایل را قبل از آپلود validate کنید
* ✔ Logging را با request ID enrich کنید
* ✔ Presigned URL مدت کوتاه داشته باشد
## 🔄 توسعه در آینده
##### این ماژول به سادگی قابل گسترش است:
- پیاده‌سازی S3 Adapter
- اضافه کردن Metadata persistence
- اضافه کردن Event-driven publishing
- اضافه کردن OpenTelemetry tracing
- پیاده‌سازی Rate limiting در adaptor
## 📌 جمع‌بندی
##### Filestore یک ماژول:
- 🧱 مستقل از فریم‌ورک
- 🧩 قابل جایگزینی
- 🚀 آماده استفاده در Production
- 🔐 امن
- ⚙️ بهینه و scalable
- 🧠 طراحی‌شده با اصول Clean Architecture
