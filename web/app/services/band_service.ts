import type Band from "~/types/band";
import type Genre from "~/types/genre";

export default class BandService {
  public GetBands(): Array<Band> {
    return new Array<Band>(
      {
        name: 'STRESSSYSTEEM',
        playing_at: '20:00',
        genres: new Array<Genre>(
          {
            name: 'Hardcore Punk',
          }
        ),
        link: 'https://open.spotify.com/artist/3kkqvvATi7otRY281oz1sO',
      },
      {
        name: 'Suicide Prophet',
        playing_at: '20:45',
        genres: new Array<Genre>(
          {
            name: 'Hardcore',
          },
          {
            name: 'Hardcore Groove',
          }
        ),
        link: 'https://open.spotify.com/artist/3Mq18rKQ2LQQKRVNkxVZTB',
      },
      {
        name: 'Razorblade Messiah',
        playing_at: '21:45',
        genres: new Array<Genre>(
          {
            name: 'Old School Thrash',
          },
          {
            name: 'Speed Metal',
          }
        ),
        link: 'https://open.spotify.com/artist/3b6qKKVYL64xsze3bdznOx',
      },
      {
        name: '2Destroy',
        playing_at: '23:00',
        genres: new Array<Genre>(
          {
            name: 'Thrash',
          },
          {
            name: 'Death Metal',
          },
        ),
        link: 'https://open.spotify.com/artist/1dH9TMeSypSqX1oJ8wGp34',
      },
    )
  }
};
