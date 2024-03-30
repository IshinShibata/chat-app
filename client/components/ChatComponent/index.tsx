"use client";
import React, { useState } from "react";

interface ChatData {
  actor: string;
  message: string;
}

export const ChatComponent: React.FC = () => {
  const [content, setContent] = useState<string>("");
  const [chatList, setChatList] = useState<string[]>([]);

  const onChangeContent = (e: React.ChangeEvent<HTMLInputElement>) => {
    setContent(e.target.value);
  };

  const postChat = async () => {
    setChatList([...chatList, content]);
    setContent("");
  };

  return (
    <div className="container mx-auto w-4/6 outline outline-2">
      <div className="bg-blue-500 text-white p-4">
        <h1 className="text-lg">チャットルーム名</h1>
      </div>
      {/* 受信者 */}
      <div className="flex flex-col space-y-2 p-4">
        <div className="bg-blue-100 text-gray-800 rounded-lg rounded-br-none shadow p-3 max-w-xs self-start">
          こんにちは、どうも！
        </div>
      </div>
      {/* 送信者 */}
      <div className="flex flex-col space-y-2 p-4">
        <div className="bg-green-100 text-gray-800 rounded-lg rounded-bl-none shadow p-3 max-w-xs self-end">
          こんにちは！
        </div>
      </div>
      <div className="p-4">
        <form className="flex gap-2 items-center">
          <input
            type="text"
            className="flex-1 p-2 border-2 border-gray-300 rounded-full focus:outline-none focus:border-blue-500 transition-colors"
            placeholder="メッセージを入力..."
          />
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-600 text-white rounded-full px-6 py-2 focus:outline-none transition-colors"
          >
            送信
          </button>
        </form>
      </div>
    </div>
  );
};
