<script lang="ts">
  import type { Goly } from "../routes/+page";
  import { goto } from "$app/navigation";
  import { feedback } from "./FeedbackStore";
  import Clipboard from "svelte-clipboard";
  import FaPaperclip from "svelte-icons/fa/FaPaperclip.svelte";
  export let data: Goly;

  let deleteGoly = false;

  function toggleDelete() {
    deleteGoly = !deleteGoly;
  }

  async function handleDelete() {
    try {
      await fetch(`/api/goly/${data.id}`, { method: "DELETE" });
      //Set timeout to make sure DB is updated
      setTimeout(() => {
        // Invalidate the page data to get a reload
        goto("/", { invalidateAll: true });
        deleteGoly = false;
      }, 100);
      feedback.set("Successfully deleted.");
    } catch (error) {
      feedback.set("An error has occurred. Please try again.");
    }
  }
</script>

<div class="rounded-lg border border-sky-500 text-sky-500 p-3">
  {#if !deleteGoly}
    <div class="flex justify-between">
      <p><span class="text-sky-900 text-bold text-lg">Goly:</span> {data.goly}</p>
      <Clipboard
        text={`https://www.goly.com/r/${data.goly}`}
        let:copy
        on:copy={() => {
          alert("Goly copied to clipboard");
        }}
      >
        <div class="cursor-pointer text-sky-900 h-4 w-4" on:click={copy} on:keypress={copy}>
          <FaPaperclip />
        </div>
      </Clipboard>
    </div>
    <p><span class="text-sky-900 text-bold text-lg">Redirect:</span> {data.redirect}</p>
    <p><span class="text-sky-900 text-bold text-lg">Clicked:</span> {data.clicked}</p>
    <div class="flex justify-between pt-3">
      <button class="bg-sky-500 px-3 py-2 rounded-md text-white"><a href={`/${data.id}`}>Update</a></button>
      <button class="bg-sky-900 px-3 py-2 rounded-md text-white" on:click={toggleDelete}>Delete</button>
    </div>
  {:else}
    <div class="h-full flex flex-col justify-center">
      <p class="text-center text-lg">Are you sure you want to delete?</p>
      <div class="flex justify-between pt-3">
        <button class="bg-sky-500 px-3 py-2 rounded-md text-white" on:click={handleDelete}>Yes</button>
        <button class="bg-sky-900 px-3 py-2 rounded-md text-white" on:click={toggleDelete}>No</button>
      </div>
    </div>
  {/if}
</div>
