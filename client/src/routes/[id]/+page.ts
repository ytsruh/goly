import type { PageLoad } from "./$types";

export type Goly = {
  clicked: number;
  goly: string;
  id: number;
  random: boolean;
  redirect: string;
};

export const load = (async ({ fetch, params }: any) => {
  try {
    const res = await fetch(`/api/goly/${params.id}`);
    const data: Goly = await res.json();

    return data;
  } catch (error) {
    console.log(error);
  }
}) satisfies PageLoad;
