# validate-credentials

Action for validate credentials for GitHub / Telegram account

## Usage

Action has input parameters:
- **github_token**
- **telegram_token** and **telegram_chat_id**

### Example: validate GitHub token

.github/workflows/credentials.yaml
```yaml
name: Validate credentials

on:
  workflow_dispatch:
  schedule:
    - cron: '00 16 */2 * *'

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - name: Validate
        uses: ci-space/validate-credentials@master
        with:
          github_token: ${{ secrets.DEP_TOKEN }}
```

### Example: validate Telegram token

.github/workflows/credentials.yaml
```yaml
name: Validate credentials

on:
  workflow_dispatch:
  schedule:
    - cron: '00 16 */2 * *'

jobs:
  sync:
    runs-on: ubuntu-latest
    steps:
      - name: Validate
        uses: ci-space/validate-credentials@master
        with:
          telegram_token: ${{ secrets.TELEGRAM_RELEASES_TOKEN }}
          telegram_chat_id: ${{ secrets.TELEGRAM_RELEASES_CHAT_ID }}
```
