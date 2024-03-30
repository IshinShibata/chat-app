"use client";
import styles from "./header.module.scss";

import { useRouter } from "next/navigation";

export const Header: React.FC = () => {
  const router = useRouter();

  return (
    <div className="headerArea">
      <button className="bg-red-500 rounded hover:bg-green-700 text-white font-bold py-2 px-4" onClick={() => router.push("/")}>
        Next.jsの勉強
      </button>
    </div>
  );
};
