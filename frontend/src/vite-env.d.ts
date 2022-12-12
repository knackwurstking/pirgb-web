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

interface SectionData {
  id: number;
  pulse?: number;
  lastPulse?: number;
  color?: number[];
  pins?: Pin[];
  element?: any;
}

type Sections = SectionData[];

interface Device {
  host: string;
  port: number;
  sections: Sections;
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
