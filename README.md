# Go Day 4 - GORM

## รันทีละบรรทัด

`go get gorm.io/gorm` // ติดตั้ง GORM ใช้เพิ่มลบข้อมูลใน Database ด้วยการโค้ด (ORM - Object Relational Mapping)

`go get gorm.io/driver/postgres` // ติดตั้ง PostgreSQL driver สำหรับ GORM เพื่อให้ GORM สามารถเชื่อมต่อกับ PostgreSQL ได้

`go get github.com/caarlos0/env/v11` // ติดตั้ง package สำหรับอ่านจากไฟล์ .env (Environment Variables)

`go install github.com/pressly/goose/v3/cmd/goose@latest` // ติดตั้ง Goose ซึ่งเป็นเครื่องมือสำหรับจัดการ Database เช่น เพิ่ม,ลบ TABLE ต่างๆ(Migration)

`goose down` // ใช้สำหรับยกเลิกการ Migration ล่าสุด (Rollback) ตามไฟล์ในโฟลเดอร์ migrations

`goose up` // ใช้สำหรับทำ Migration ล่าสุด (เช่น สร้าง TABLE ใหม่) ตามไฟล์ในโฟลเดอร์ migrations

`go run cmd/student/main.go` // รันคำสั่งซึ่งจะเชื่อมต่อกับ Database และทำงานตามที่เรากำหนดไว้ในโค้ด
