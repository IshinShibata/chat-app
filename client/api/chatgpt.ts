export interface PostData {
  message: string | null;
  sessionId: string | null;
}

export const postChat = async (data: PostData) => {
  const sample = {
    message: "こんにちは",
    sessionId: "1234",
  };
  return sample;
};
