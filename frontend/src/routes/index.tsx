import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import RegisterPage from '@/pages/Register';
// 今後 Login や Home ができたらここに追加していく

const router = createBrowserRouter([
  {
    path: '/register',
    element: <RegisterPage />,
  },
  {
    /* ログイン画面（今後作成） */
    path: '/login',
    element: <div>ログイン画面（作成予定）</div>,
  },
  {
    /* セットアップ画面（今後作成） */
    path: '/setup',
    children: [
      { path: 'name', element: <div>名前設定画面（作成予定）</div> },
      { path: 'image', element: <div>画像設定画面（作成予定）</div> },
    ],
  },
]);

export const AppRouter = () => {
  return <RouterProvider router={router} />;
};