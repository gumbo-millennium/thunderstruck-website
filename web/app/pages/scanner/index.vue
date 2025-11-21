<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    {{ ticket?.value }} {{ scanned }} {{ paused }}
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />
      <main
        v-if="authenticated"
        class="flex flex-col gap-4 p-4"
      >
        <div
          v-if="ticket"
          class="absolute w-screen h-screen flex flex-col justify-center items-center left-0 top-0 pb-32 z-100 text-6xl text-white font-archivo text-center"
          :class="{
            'bg-green-primary': (ticket?.type === TicketType.ENTRY) && (ticket?.state !== TicketState.USED),
            'bg-purple': (ticket?.type === TicketType.CREW) && (ticket?.state !== TicketState.USED),
            'bg-orange-500': ticket?.state === TicketState.USED,
            'bg-red-500': !ticket,
          }"
          @click.stop="clearStatus"
        >
          {{ ticket?.type }} - {{ ticket?.state }}
          <p
            v-if="ticket?.state === TicketState.USED"
            class="text-xl"
          >Check voor polsbandje</p>
        </div>
        <div
          v-if="scanned && !ticket"
          class="absolute w-screen h-screen flex flex-col justify-center items-center bg-red-500 left-0 top-0 pb-32 z-100 text-6xl text-white font-archivo"
          @click.stop="clearStatus"
        >
          X
        </div>
        <LargeHeaderComponent>
          Ticket scanner
        </LargeHeaderComponent>
        <p
          v-if="error"
          class="rounded color-white p-2 bg-red-300 outline-2 outline-red-500"
        >
          {{ error }}
        </p>
        <div
          class="relative flex justify-center items-center h-96 w-full"
        >
          <h1
            class="absolute text-red-300 text-3xl z-50"
          >
            
            <img
              v-if="paused"
              class="w-64 aspect-square animate-pulse"
              src="~/assets/images/gumbo.webp"
            >
          </h1>
          <QrcodeStream
            :track="paintBoundingBox"
            :paused="paused"
            @error="onError"
            @detect="onDetect"
          />
        </div>
      </main>
      <main
        v-else
        class="flex flex-col gap-4 p-4"
      >
        <LargeHeaderComponent>
          Ticket scanner
        </LargeHeaderComponent>
        <p>
          Voer de sleutel in om door te gaan.
        </p>
        <p
          v-if="error"
          class="rounded color-white p-2 bg-red-300 outline-2 outline-red-500"
        >
          {{ error }}
        </p>
        <input
          v-model="token"
          placeholder="Sleutel"
          required="true"
          class="outline-purple outline-2 bg-purple-100 rounded p-2"
          type="password"
          @keyup.enter="validateToken"
        >
        <CallToActionComponent
          :disabled="disabled"
          @click="validateToken"
        >
          Login
        </CallToActionComponent>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { QrcodeStream, type DetectedBarcode } from 'vue-qrcode-reader';
import type Ticket from '~/types/ticket';
import { TicketState, TicketType } from '~/types/ticket';

const token: Ref<string> = ref('');
const disabled: Ref<boolean> = ref(true);
const error: Ref<string> = ref('');
const paused: Ref<boolean> = ref(false);
const authenticated: Ref<boolean> = ref(false); // Insecure authentication but all I'm gonna do for now

const scanned: Ref<boolean> = ref(false);
const ticket: Ref<Ticket | undefined> = ref();
const statusTimeout: Ref<number> = ref(0);
const pauseTimeout: Ref<number> = ref(0);

watch(token, (to) => {
  disabled.value = to == '';
});

 function paintBoundingBox(detectedCodes, ctx) {
  for (const detectedCode of detectedCodes) {
    const {
      boundingBox: { x, y, width, height }
    } = detectedCode;

    ctx.lineWidth = 2;
    ctx.strokeStyle = '#007bff';
    ctx.strokeRect(x, y, width, height);
  }
}

function onError(err: unknown) {
  error.value = `[${err.name}]: `;

  if (err.name === 'NotAllowedError') {
    error.value = 'Geef de pagina toestemming om je camera te gebruiken';
  } else if (err.name === 'NotFoundError') {
    error.value = 'Dit apparaat ondersteunt geen camera';
  } else if (err.name === 'NotReadableError') {
    error.value = 'Is de camera al in gebruik?';
  } else {
    error.value += err.message
  }
}

async function onDetect(detected: Array<DetectedBarcode>) {
  statusTimeout.value = setTimeout(() => {
    clearStatus();
  }, 3000);

  if (detected.length === 0) return;
  
  const value: string = detected[0]!.rawValue;

  try {
    const response: Ticket = await useClientFetch<Ticket>('scanner', {
      method: 'POST',
      body: {
        token: token.value,
        ticket: value,
      },
    });

    ticket.value = response;

  } catch(e: unknown) {
    if (e.data.error === 'token may not be empty') {
      authenticated.value = false;
      return;
    }

    if (e.data.error !== 'no rows in result set') {
      error.value = e.data.error;
    }
  }
  scanned.value = true;
}

function clearStatus() {
  paused.value = true;
  pauseTimeout.value = setTimeout(() => {
    clearTimeout(pauseTimeout.value);
    paused.value = false;
  }, 500);

  ticket.value = undefined;
  scanned.value = false;
  clearTimeout(statusTimeout.value);
}

async function validateToken() {
  error.value = '';

  if (token.value === '') {
    error.value = 'Sleutel mag niet leeg zijn.';
    return;
  }

  try {
    await useClientFetch('scanner/token', {
      method: 'POST',
      body: {
        token: token.value,
      },
    })

    authenticated.value = true;

  } catch(e: unknown) {
    error.value = 'Incorrecte sleutel ingevoerd.';
  }
}
</script>
