import { useState, useEffect } from 'react';
import { Shield, Zap, Terminal, Activity, Loader2 } from 'lucide-react';
import axios from 'axios';

interface Entity {
  id: string;
  name: string;
  status: string;
  metadata: string;
}

const API_BASE = 'http://localhost:8080';

function App() {
  const [entities, setEntities] = useState<Entity[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const initBunker = async () => {
      try {
        setLoading(true);
        
        // 1. Obtener Token (Simulado en este hito)
        const authRes = await axios.get(`${API_BASE}/login`);
        const token = authRes.data.token;

        // 2. Traer Entidades Reales
        const entitiesRes = await axios.get(`${API_BASE}/api/entities`, {
          headers: { Authorization: `Bearer ${token}` }
        });

        setEntities(entitiesRes.data || []);
        setError(null);
      } catch (err) {
        console.error('Bunker connection failed', err);
        setError('OFFLINE - Check Go API and Docker DB');
        
        // Fallback para dev sin API prendida
        setEntities([
          { id: '1', name: 'Hermes', status: 'Core Active', metadata: '' },
          { id: '2', name: 'Decepticon', status: 'Red Team Waiting', metadata: '' },
        ]);
      } finally {
        setLoading(false);
      }
    };

    initBunker();
  }, []);

  return (
    <div className="min-h-screen w-full flex flex-col items-center justify-center p-8 bg-bunker-bg">
      {/* Header */}
      <header className="mb-12 text-center">
        <h1 className="text-6xl font-bold text-primary mb-2 tracking-tighter text-glow-cyan uppercase italic">
          Bunker OS
        </h1>
        <p className="text-secondary font-mono tracking-widest uppercase text-sm animate-pulse flex items-center justify-center gap-2">
          {loading && <Loader2 size={14} className="animate-spin" />}
          Agentic Knowledge Vault — v1.4.1
        </p>
        {error && <p className="text-red-500 text-[10px] mt-2 font-mono uppercase tracking-widest">{error}</p>}
      </header>

      {/* Main Grid */}
      <main className="grid grid-cols-1 md:grid-cols-2 gap-6 max-w-4xl w-full">
        {/* API Status Card */}
        <div className="glass p-6 rounded-xl glow-cyan group hover:scale-105 transition-transform cursor-pointer">
          <div className="flex items-center justify-between mb-4">
            <Activity className={error ? "text-red-500" : "text-primary"} size={32} />
            <span className={`${error ? "bg-red-500/20 text-red-400 border-red-500/30" : "bg-green-500/20 text-green-400 border-green-500/30"} text-xs px-2 py-1 rounded-full border`}>
              {error ? "OFFLINE" : "ONLINE"}
            </span>
          </div>
          <h2 className="text-xl font-bold mb-1">OZY-API</h2>
          <p className="text-white/60 text-sm font-mono">{API_BASE}/health</p>
        </div>

        {/* Security Audit Card */}
        <div className="glass p-6 rounded-xl glow-purple group hover:scale-105 transition-transform cursor-pointer">
          <div className="flex items-center justify-between mb-4">
            <Shield className="text-secondary" size={32} />
            <span className="bg-purple-500/20 text-purple-400 text-xs px-2 py-1 rounded-full border border-purple-500/30">SECURE</span>
          </div>
          <h2 className="text-xl font-bold mb-1">TRIVY AUDIT</h2>
          <p className="text-white/60 text-sm font-mono">Last scan: Clean</p>
        </div>

        {/* Entities List (Real Data) */}
        <div className="glass p-6 rounded-xl md:col-span-2">
          <div className="flex items-center gap-2 mb-6">
            <Terminal className="text-primary" size={24} />
            <h2 className="text-lg font-bold uppercase tracking-widest">Live Entities Registry</h2>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
            {entities.map((e) => (
              <div key={e.id} className="flex flex-col items-center p-3 rounded-lg bg-white/5 border border-white/10 hover:border-primary/50 transition-colors">
                <span className="font-bold text-white uppercase text-xs">{e.name}</span>
                <span className="text-[10px] text-primary uppercase tracking-tighter mt-1">{e.status}</span>
              </div>
            ))}
            {entities.length === 0 && !loading && (
              <p className="col-span-full text-white/20 text-xs font-mono text-center">No entities found in database.</p>
            )}
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="mt-12 flex items-center gap-4 text-white/30 font-mono text-xs uppercase tracking-[0.3em]">
        <Zap size={14} className="text-primary" />
        Bunker connected via n8n & OpenCode
      </footer>
    </div>
  );
}

export default App;
