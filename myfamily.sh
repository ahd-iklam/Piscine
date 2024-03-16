if [ -z "$HERO_ID" ]; then
    echo "HERO_ID environment variable is not set."
    exit 1
fi
json=$(curl -s "https://learn.zone01oujda.ma/assets/superhero/all.json")
family=$(echo "$json" | jq  --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id|tonumber)) | .connections.relatives')
formatted_family=$(echo "$family" | tr -d '"')
echo "$formatted_family"