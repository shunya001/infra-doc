infra-doc
Infrastructure document generator CLI written in Go.

インフラ設計書を自動生成する Go 製 CLI ツールです。

Features
Scan Linux server information
Export system information as JSON
Import JSON data
Generate infrastructure documents
Generate parameter sheets
Generate build procedures
Collected Information
Hostname
OS
CPU
Memory
Disk
Network Interface
IP Address
Installation
git clone https://github.com/shunya001/infra-doc.git

cd infra-doc

go mod tidy
Usage
Scan Linux Server
go run . scan -p web-prod
Generated files:

output/
├── server_info.json
├── system_info.md
├── basic_design.md
└── parameter_sheet.md
Generate Documents from JSON
go run . import output/server_info.json
Generated files:

output/
├── basic_design.md
├── parameter_sheet.md
└── system_info.md
Example JSON
{
  "Project": "web-prod",
  "Hostname": "server01",
  "OS": "Ubuntu 24.04",
  "CPU": "Intel Xeon",
  "Memory": "8Gi",
  "Disk": "50G",
  "NetworkInterface": "eth0",
  "IPAddress": "192.168.1.10"
}
Roadmap
Markdown document generation
Linux system information collection
JSON export
JSON import
Excel export (.xlsx)
Multi-server inventory support
Terraform integration
AWS integration
日本語
概要
Linuxサーバの情報を取得し、設計書やパラメータシートを自動生成するCLIツールです。

主な機能
サーバ情報取得
JSON出力
JSON読込
基本設計書生成
パラメータシート生成
構築手順書生成
取得情報
ホスト名
OS
CPU
メモリ
ディスク
NIC
IPアドレス
今後予定
Excel出力
複数サーバ対応
Terraform連携
AWS連携