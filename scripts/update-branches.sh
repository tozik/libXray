#!/bin/bash

# Скрипт для получения актуального списка веток из репозитория tozik/libXray
# и обновления workflow файла

echo "Получение списка веток из репозитория tozik/libXray..."

# Получаем ветки через GitHub API (требует токен в переменной GITHUB_TOKEN)
if [ -z "$GITHUB_TOKEN" ]; then
    echo "Используем публичный API (лимит 60 запросов в час)"
    BRANCHES=$(curl -s "https://api.github.com/repos/tozik/libXray/branches" | jq -r '.[].name' | head -10)
else
    echo "Используем аутентифицированный API"
    BRANCHES=$(curl -s -H "Authorization: token $GITHUB_TOKEN" \
        "https://api.github.com/repos/tozik/libXray/branches" | jq -r '.[].name' | head -10)
fi

echo "Найденные ветки:"
echo "$BRANCHES"

echo ""
echo "Для обновления workflow файла, замените секцию options на:"
echo "        options:"
while IFS= read -r branch; do
    echo "          - '$branch'"
done <<< "$BRANCHES"
echo "          - 'custom'"
