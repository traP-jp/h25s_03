import { HttpResponse } from "msw";
import { createOpenApiHttp } from "openapi-msw";
import type { paths } from "@/api/schema";

export const http = createOpenApiHttp<paths>();
export const handlers = [
  http.get("/events", () => {
    return HttpResponse.json([
      {
        event_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
        title: "Event 1",
        description: "Description for Event 1",
        date: "2023-10-01",
        is_open: true,
        admins: [],
        is_me_attendee: false,
      },
    ]);
  }),
  http.get("/events/{eventID}", () => {
    return HttpResponse.json({
      event_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      title: "忘年会",
      description: "楽しいイベント",
      date: "2025-06-21T06:06:28.642Z",
      is_open: true,
      admins: ["miyamon", "ogu_kazemiya"],
      attendees: [
        "miyamon",
        "ogu_kazemiya",
        "ten_ten",
        "cp20",
        "Eraxyso",
        "hachimitsu",
        "Naph1",
      ],
      created_at: "2025-06-21T06:06:28.642Z",
      updated_at: "2025-06-21T06:06:28.642Z",
      is_deleted: false,
      is_me_attendee: true,
    });
  }),
  http.get("/events/{eventID}/lotteries/{lotteryID}", () => {
    return HttpResponse.json({
      lottery_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      event_id: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      title: "焼肉食べ放題",
      is_deleted: false,
      created_at: "2025-06-21T06:06:28.642Z",
      updated_at: "2025-06-21T06:06:28.642Z",
      winners: ["miyamon", "ogu_kazemiya", "ten_ten"],
    });
  }),
  http.post("/events/{eventID}/lotteries/{lotteryID}", () => {
    return HttpResponse.json({
      winner: "miyamon",
    });
  }),
];
