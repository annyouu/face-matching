FROM node:20-alpine

WORKDIR /app

# package.jsonを先にコピー（キャッシュ利用のため）
COPY package.json package-lock.json ./
RUN npm install

# 全ファイルをコピー
COPY . .

# 開発用サーバーを起動
CMD ["npm", "run", "dev"]