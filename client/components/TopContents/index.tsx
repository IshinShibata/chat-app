import { useRouter } from "next/navigation";
export const TopContents: React.FC = () => {
  const router = useRouter();
  return (
    <div className="container">
      <h1 className="pageTitle">TOPページ</h1>
      <div className="box">
        <div>
          <p className="userName">ユーザー名：shibata</p>
          <div>
            <button
              className="bg-green-500 rounded hover:bg-green-700 text-white font-bold py-2 px-4"
              onClick={() => {
                router.push("/chat");
              }}
            >
              Chatページへ
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
