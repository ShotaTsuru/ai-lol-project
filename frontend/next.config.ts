import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Turbopackの設定（stable）
  turbopack: {
    rules: {
      '*.svg': {
        loaders: ['@svgr/webpack'],
        as: '*.js',
      },
    },
  },
  // 開発環境でのファストリフレッシュを有効化
  reactStrictMode: true,
};

export default nextConfig;
