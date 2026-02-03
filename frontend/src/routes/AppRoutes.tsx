import { Routes, Route } from 'react-router-dom';

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<div>ホーム画面（開発中）</div>} />
            <Route path="/login" element={<div>ログイン画面（開発中）</div>} />
            <Route path="/register" element={<div>新規登録画面（開発中）</div>} />
            <Route path="/inventory" element={<div>在庫一覧（開発中）</div>} />
        </Routes>
    );
};

export default AppRoutes;