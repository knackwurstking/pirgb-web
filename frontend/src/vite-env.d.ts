/// <reference types="svelte" />
/// <reference types="vite/client" />

interface Section {
  id: number;
  pulse?: number;
  lastPulse?: number;
  // TODO: ...
}

interface Device {
  host: string;
  port: number;
  sections: Section[];
  groups: string[];
}

type Devices = Device[];
