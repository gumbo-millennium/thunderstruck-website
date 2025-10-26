<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />
      <main class="flex flex-col gap-10 p-4">
        <CallToActionComponent
          class="hover:animate-pulse"
          @click="() => console.log('hi')"
        >
          Grijp tickets
        </CallToActionComponent>

        <div class="flex justify-between uppercase">
          <div class="flex flex-col leading-4 w-36">
            <h1 class="font-archivo text-purple">Wanneer</h1>
            <h2>21 november, <span class="font-black"><b>20:00</b></span></h2>
          </div>
          <div class="flex flex-col leading-4 text-right w-36">
            <h1 class="font-archivo text-purple">Waar</h1>
            <h2>Het Vliegende Paard, <span class="font-black"><b>Zwolle</b></span></h2>
          </div>
        </div>

        <div class="flex flex-col gap-8">
          <LargeHeaderComponent>
            Programma
          </LargeHeaderComponent>

          <ProgramItemComponent
            v-for="band, i in bands"
            :key="i"
            :time="band.playing_at"
            :title="band.name"
            :genres="band.genres"
            :link="band.link"
          />
        </div>

        <div>
          <LargeHeaderComponent>
            Onze partners
          </LargeHeaderComponent>

          <PartnerCarousal
            class="mt-4"
            :partners="partners"
          />
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import PartnerCarousal from '~/components/PartnerCarousal.vue';
import BandService from '~/services/band_service';
import PartnerService from '~/services/partner_service';
import type Band from '~/types/band';
import type Partner from '~/types/partner';

const bandService: BandService = new BandService();
const partnerService: PartnerService = new PartnerService();

const bands: Ref<Array<Band>> = ref(bandService.GetBands());
const partners: Ref<Array<Partner>> = ref(partnerService.GetPartners());
</script>
