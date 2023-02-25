import type { PageLoad } from "./$types";

export type Goly = {
  clicked: number;
  goly: string;
  id: number;
  random: boolean;
  redirect: string;
};

export const load = (async ({ fetch }: any) => {
  try {
    const res = await fetch("/api/goly");

    const data: Goly[] = await res.json();

    return { golies: data };
  } catch (error) {
    console.log(error);
  }
}) satisfies PageLoad;
