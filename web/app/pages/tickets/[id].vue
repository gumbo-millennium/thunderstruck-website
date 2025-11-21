<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />
      <main class="flex flex-col gap-4 p-4">
        <LargeHeaderComponent>
          Je ticket
        </LargeHeaderComponent>
        <div
          class="flex flex-col justify-center items-center"
        >
          <img
            v-if="data === undefined"
            class="animate-pulse aspect-square h-14"
            src="~/assets/images/gumbo.webp"
          >
          <ClientOnly>
            <img
              v-if="data"
              :src="qr"
              :alt="data.value"
              class="h-48 aspect-square"
            >
          </ClientOnly>
          <p
            class="text-xs"
          >{{data?.value}}</p>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import type Ticket from '~/types/ticket';
import { useQRCode } from "@vueuse/integrations/useQRCode";

const route = useRoute();
const { data, error } = await useApi<Ticket>(`tickets/${route.params.id}`)
const qr = useQRCode(data.value?.value?.toString() ?? '');

if (error.value !== undefined) {
  console.error(error.value);
  navigateTo('/');
}
</script>
