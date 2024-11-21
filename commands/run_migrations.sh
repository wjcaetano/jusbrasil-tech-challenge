#!/usr/bin/env bash
echo "Running migrations..."
mysql -h testlocal -u root -proot testlocal < /resources/sql/migrations.sql
echo "Migrations completed!"