import type Partner from "~/types/partner";

export default class PartnerService {
  public GetPartners(): Array<Partner> {
    return new Array<Partner>(
      {
        name: 'Gumbo Millennium',
        image: 'images/partners/gumbo.webp',
        link: 'https://gumbo-millennium.nl/',
      },
      {
        name: 'Buro Ruis',
        image: 'images/partners/buroruis.webp',
        link: 'https://buroruis.nl/',
      },
      {
        name: 'Het Vliegende Paard',
        image: 'images/partners/hetvliegendepaard.webp',
        link: 'https://hetvliegendepaard.nl/',
      },
      {
        name: 'SOOZ',
        image: 'images/partners/sooz.webp',
        link: 'https://sooz.nl/',
      },
      {
        name: 'Gemeente Zwolle',
        image: 'images/partners/zwolle.webp',
        link: 'https://zwolle.nl/',
      },
    );
  }
};
