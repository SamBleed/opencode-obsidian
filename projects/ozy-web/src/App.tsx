import { Shield, Zap, Terminal, Activity } from 'lucide-react';

function App() {
  const entities = [
    { name: 'Hermes', status: 'Active', color: 'text-purple-400' },
    { name: 'Decepticon', status: 'Standby', color: 'text-red-400' },
    { name: 'DeepTutor', status: 'Learning', color: 'text-cyan-400' },
    { name: 'Goose', status: 'Connected', color: 'text-green-400' },
  ];

  return (
    <div className="min-h-screen w-full flex flex-col items-center justify-center p-8 bg-bunker-bg">
      {/* Header */}
      <header className="mb-12 text-center">
        <h1 className="text-6xl font-bold text-primary mb-2 tracking-tighter text-glow-cyan uppercase italic">
          Bunker OS
        </h1>
        <p className="text-secondary font-mono tracking-widest uppercase text-sm animate-pulse">
          Agentic Knowledge Vault — v1.4.1
        </p>
      </header>

      {/* Main Grid */}
      <main className="grid grid-cols-1 md:grid-cols-2 gap-6 max-w-4xl w-full">
        {/* API Status Card */}
        <div className="glass p-6 rounded-xl glow-cyan group hover:scale-105 transition-transform cursor-pointer">
          <div className="flex items-center justify-between mb-4">
            <Activity className="text-primary" size={32} />
            <span className="bg-green-500/20 text-green-400 text-xs px-2 py-1 rounded-full border border-green-500/30">ONLINE</span>
          </div>
          <h2 className="text-xl font-bold mb-1">OZY-API</h2>
          <p className="text-white/60 text-sm font-mono">http://localhost:8080/health</p>
        </div>

        {/* Security Audit Card */}
        <div className="glass p-6 rounded-xl glow-purple group hover:scale-105 transition-transform cursor-pointer">
          <div className="flex items-center justify-between mb-4">
            <Shield className="text-secondary" size={32} />
            <span className="bg-purple-500/20 text-purple-400 text-xs px-2 py-1 rounded-full border border-purple-500/30">SECURE</span>
          </div>
          <h2 className="text-xl font-bold mb-1">TRIVY AUDIT</h2>
          <p className="text-white/60 text-sm font-mono">Last scan: 0 Hallucinations</p>
        </div>

        {/* Entities List */}
        <div className="glass p-6 rounded-xl md:col-span-2">
          <div className="flex items-center gap-2 mb-6">
            <Terminal className="text-primary" size={24} />
            <h2 className="text-lg font-bold uppercase tracking-widest">Active Entities</h2>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
            {entities.map((e) => (
              <div key={e.name} className="flex flex-col items-center p-3 rounded-lg bg-white/5 border border-white/10">
                <span className={`font-bold ${e.color}`}>{e.name}</span>
                <span className="text-[10px] text-white/40 uppercase tracking-tighter">{e.status}</span>
              </div>
            ))}
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="mt-12 flex items-center gap-4 text-white/30 font-mono text-xs uppercase tracking-[0.3em]">
        <Zap size={14} className="text-primary" />
        Built with Arch Linux + OpenCode + React 19
      </footer>
    </div>
  );
}

export default App;
