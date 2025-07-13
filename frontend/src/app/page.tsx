'use client';

import { useState } from 'react';
import { Search, BarChart3, Trophy, Users, ArrowRight, Sparkles } from 'lucide-react';

export default function Home() {
  const [searchQuery, setSearchQuery] = useState('');

  const features = [
    {
      icon: <BarChart3 className="w-8 h-8 text-[--primary]" />,
      title: "メタ分析",
      description: "AIによる最新メタゲームの深度分析と予測"
    },
    {
      icon: <Trophy className="w-8 h-8 text-[--primary]" />,
      title: "チャンピオン統計",
      description: "勝率、ピック率、バン率の詳細データと推奨ビルド"
    },
    {
      icon: <Users className="w-8 h-8 text-[--primary]" />,
      title: "プロシーン分析",
      description: "プロプレイヤーの戦略とチーム構成の分析"
    }
  ];

  return (
    <div className="min-h-screen hero-gradient">
      {/* Header */}
      <header className="bg-[--card-bg]/20">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <div className="w-8 h-8 bg-gradient-to-br from-[--primary] to-[--accent] rounded"></div>
              <span className="text-xl font-bold gradient-text">LoL Meta AI</span>
            </div>
            <nav className="hidden md:flex space-x-6">
              <a href="#" className="text-[--foreground] hover:text-[--primary] transition-colors font-medium">分析</a>
              <a href="#" className="text-[--foreground] hover:text-[--primary] transition-colors font-medium">チャンピオン</a>
              <a href="#" className="text-[--foreground] hover:text-[--primary] transition-colors font-medium">プロシーン</a>
              <a href="#" className="text-[--foreground] hover:text-[--primary] transition-colors font-medium">ランキング</a>
            </nav>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="container mx-auto px-4 py-16">
        {/* Hero Section */}
        <div className="text-center mb-16">
          <div className="mb-6">
            <h1 className="text-5xl md:text-7xl font-bold mb-4">
              <span className="gradient-text">LoL Meta</span>
              <br />
              <span className="text-[--foreground]">AI分析プラットフォーム</span>
            </h1>
            <p className="text-xl text-[--text-muted] max-w-2xl mx-auto leading-relaxed">
              League of LegendsのメタゲームをAIで分析。<br />
              チャンピオン統計、ビルド推奨、戦略分析を提供する<br />
              インテリジェントプラットフォーム
            </p>
          </div>

          {/* Search Bar */}
          <div className="max-w-2xl mx-auto mb-12">
            <div className="relative">
              <Search className="absolute left-4 top-1/2 transform -translate-y-1/2 text-[--foreground]/40 w-5 h-5" />
              <input
                type="text"
                placeholder="サモナー名、チャンピオン名で検索..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full pl-12 pr-16 py-4 bg-[--card-bg] text-[--foreground] placeholder-[--text-muted] focus:outline-none focus:bg-[--card-bg]/80 transition-all rounded-full"
              />
              <button className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-[--primary] hover:bg-[--primary]/80 text-[--background] px-6 py-2 rounded-full transition-all font-medium shadow-lg">
                検索
              </button>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-16">
            <button className="group bg-[--primary] hover:bg-[--primary]/90 text-[--background] px-8 py-3 rounded-full font-semibold transition-all flex items-center space-x-2 shadow-lg hover:shadow-xl hover:scale-105">
              <Sparkles className="w-5 h-5" />
              <span>AI分析を開始</span>
              <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </button>
            <button className="bg-[--card-bg] hover:bg-[--card-bg]/80 text-[--foreground] px-8 py-3 rounded-full font-semibold transition-all hover:scale-105">
              デモを見る
            </button>
          </div>
        </div>

        {/* Features Section */}
        <div className="grid md:grid-cols-3 gap-8 mb-16">
          {features.map((feature, index) => (
            <div key={index} className="bg-[--card-bg]/60 rounded-2xl p-8 text-center hover:bg-[--card-bg]/80 transition-all duration-300 hover:scale-105">
              <div className="mb-6 flex justify-center">
                {feature.icon}
              </div>
              <h3 className="text-xl font-bold text-[--foreground] mb-4">{feature.title}</h3>
              <p className="text-[--text-muted] leading-relaxed">{feature.description}</p>
            </div>
          ))}
        </div>

        {/* Stats Section */}
        <div className="bg-[--card-bg]/40 rounded-2xl p-8">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
            <div>
              <div className="text-3xl font-bold gradient-text mb-2">160+</div>
              <div className="text-[--text-muted]">チャンピオン</div>
            </div>
            <div>
              <div className="text-3xl font-bold gradient-text mb-2">1M+</div>
              <div className="text-[--text-muted]">試合データ</div>
            </div>
            <div>
              <div className="text-3xl font-bold gradient-text mb-2">99.2%</div>
              <div className="text-[--text-muted]">予測精度</div>
            </div>
            <div>
              <div className="text-3xl font-bold gradient-text mb-2">24/7</div>
              <div className="text-[--text-muted]">リアルタイム更新</div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="backdrop-blur-sm mt-16">
        <div className="container mx-auto px-4 py-8">
          <div className="text-center text-[--text-muted]">
            <p>&copy; 2024 LoL Meta AI. League of Legends AI分析プラットフォーム</p>
            <p className="text-sm mt-2">Riot Games公式ではありません。League of LegendsはRiot Games, Inc.の商標です。</p>
          </div>
        </div>
      </footer>
    </div>
  );
}
