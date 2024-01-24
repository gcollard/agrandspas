<p align="center">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="./logo-AGP.png">
      <img alt="a grands pas logo" src="./logo-AGP.png">
    </picture>
</p>

---

# À grands pas

"À petits pas" missing functionality to bulk download and backup all pictures and videos.

### Who is this for?
1. Busy parents
2. Day care center staff 

## How to use?
```bash
# Replace with your "À petits pas" account credentials
AGP_MYID=accountid AGP_USERNAME=your@email.com AGP_PASSWORD=password go run .
```

## What will it do?
"A grands pas" will start downloading every media available to your account.
1. **All types of medias**: Every pictures & videos will be saved.
2. **Archive mode**: The last 366 days of media will be downloaded.
3. **Multiple kids support**: If multiple kids are attached to your account, all their photos will be downloaded as well.

### Settings 
Settings are updated using environment variables.
| Variable | Default Value | Description |
| -------- | ------------- | ----------- |
| AGP_DEST_FOLDER | `./media` | Path to destination folder |
| AGP_MYID |  | Account PersonID |
| AGP_USERNAME |  | Account ID (Email address) |
| AGP_PASSWORD |  | Account Password |
 

### Caveats
1. Only the email connection method is supported.

### Roadmap
1. Find account id automatically
2. Build binary package
3. web & mobile app