import type Genre from "./genre";

export default interface Band {
  name: string,
  playing_at: string,
  genres: Array<Genre>,
  link?: string,
};
