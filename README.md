<p align="center">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="./logo-AGP.png">
      <img alt="a grands pas logo" src="./logo-AGP.png">
    </picture>
</p>

---

# À grands pas

"À petits pas" missing functionality to bulk download and backup your account pictures and videos.

### Who is this for?
1. Busy parents
2. Day care center staff 

## How to use?
```bash
# Replace with your "À petits pas" account credentials
AGP_USERNAME=your@email.com AGP_PASSWORD=password go run .
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
| AGP_USERNAME |  | Account ID (Email address) |
| AGP_PASSWORD |  | Account Password |
| AGP_DEST_FOLDER | `./Downloads` | Path to destination folder |
| AGP_DAYS_TO_FETCH | `366` | Number of past days to fetch (between 1 and 366) |

 

### Caveats
1. No support for social auth. Only email authentification is supported.

### Roadmap
- [x] Find account id automatically
- [x] Build binary package
- [ ] Web & mobile app
