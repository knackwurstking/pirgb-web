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
  pins?: Pin[];
}

interface Device {
  host: string;
  port: number;
  sections: Section[];
  groups: string[];
}

type Devices = Device[];

type BaseEventTypes = "offline" | "online" | "change";

interface BaseEvent<N, D> {
  name: N;
  data: D;
}

interface DeviceEvent extends BaseEvent {
  host: string;
  port: number;
}

interface ChangeEvent extends BaseEvent {
  host: string;
  port: number;
  id: number;
  pulse?: number;
  lastPulse?: number;
  color?: number[];
  pins?: Pin[];
}
