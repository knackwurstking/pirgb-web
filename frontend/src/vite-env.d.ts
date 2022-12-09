/// <reference types="svelte" />
/// <reference types="vite/client" />

interface Pin {
    pin: number; 
    range: number;
    pulse: number;
    colorValue: number;
    colorPulse: number;
    isRunning: boolean;
}

interface Section {
  id: number;
  pulse?: number;
  lastPulse?: number;
  color?: number[];
  pins?: SectionPin[];
}

interface Device {
  host: string;
  port: number;
  sections: Section[];
  groups: string[];
}

type Devices = Device[];
