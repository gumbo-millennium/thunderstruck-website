<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />

      <!-- Desktop -->
      <main
        class="hidden lg:flex flex-col px-4 pb-4 w-full gap-4"
      >
        <div class="w-200 mx-auto relative">
          <div class="absolute flex justify-center items-center w-full bg-white rounded-2xl left-0 top-[-10rem] p-4">
            <div class="w-full flex justify-between gap-4 uppercase">
              <div class="flex flex-col leading-4 w-36">
                <h1 class="font-archivo text-purple">Wanneer</h1>
                <h2>21 november, <span class="font-black"><b>20:00</b></span></h2>
              </div>

              <CallToActionComponent
                class="hover:animate-pulse"
                @click="() => console.log('hi')"
              >
                Grijp tickets
              </CallToActionComponent>

              <div class="flex flex-col leading-4 w-36 text-right">
                <h1 class="font-archivo text-purple">Waar</h1>
                <h2>Het Vliegende Paard, <span class="font-black"><b>Zwolle</b></span></h2>
              </div>
            </div>
          </div>
        </div>

        <div class="flex flex-col gap-4">
          <LargeHeaderComponent>
            Programma
          </LargeHeaderComponent>

          <div class="flex justify-between gap-4">
            <ProgramItemComponent
              v-for="band, i in bands"
              :key="i"
              class="w-1/4"
              :time="band.playing_at"
              :title="band.name"
              :genres="band.genres"
              :link="band.link"
            />
          </div>
        </div>

        <div class="flex flex-col gap-4">
          <LargeHeaderComponent>
            Impressie
          </LargeHeaderComponent>

          <div class="flex justify-between gap-4">
            <img
              class="rounded-2xl w-1/3"
              src="/images/impression/impression_01.webp"
            >
            <img
              class="rounded-2xl w-1/3"
              src="/images/impression/impression_02.webp"
            >
            <img
              class="rounded-2xl w-1/3"
              src="/images/impression/impression_03.webp"
            >
          </div>
        </div>

        <div class="flex flex-col gap-4">
          <LargeHeaderComponent>
            Onze partners
          </LargeHeaderComponent>

          <div class="w-full flex justify-between xl:justify-start gap-12">
            <a
              v-for="partner, i in partners"
              :key="i"
              ref="noreferrer noopener"
              target="_blank"
              :href="partner.link"
              :title="partner.name"
              class="flex justify-center items-center w-24 aspect-square"
            >
              <img
                :src="partner.image"
                :alt="partner.name"
              >
            </a>
          </div>
        </div>
      </main>

      <!-- Mobile -->
      <main
        v-if="isMobile"
        class="lg:hidden flex flex-col gap-10 p-4"
      >
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

    <FloatingCallToActionComponent
      v-if="!isMobile"
    >
      Grijp tickets
    </FloatingCallToActionComponent>
  </div>
</template>

<script setup lang="ts">
import PartnerCarousal from '~/components/PartnerCarousal.vue';
import BandService from '~/services/band_service';
import PartnerService from '~/services/partner_service';
import type Band from '~/types/band';
import type Partner from '~/types/partner';

onMounted(() => {
  isMobile.value = window.innerWidth < 1024;

  window.addEventListener('resize', () => {
    isMobile.value = window.innerWidth < 1024;
  });
});

const bandService: BandService = new BandService();
const partnerService: PartnerService = new PartnerService();

const bands: Ref<Array<Band>> = ref(bandService.GetBands());
const partners: Ref<Array<Partner>> = ref(partnerService.GetPartners());
const isMobile: Ref<boolean> = ref(false);
</script>
