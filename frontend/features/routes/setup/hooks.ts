// APIを叩いた後の画面遷移ロジックを書く役割
import { useRouter } from "next/navigation";
import { setupEndpoints } from "./endpoint";
import { UserStatus } from "@/type/user";

export const useSetup = () => {
    const router = useRouter();

    const navigateByStatus = (status: UserStatus) => {
        switch (status) {
            case "ACTIVE":
                router.push("/home");
                break;
            case "PENDING_IMAGE":
                router.push("/setup/image");
                break;
            case "PENDING_NAME":
                router.push("/setup/name");
                break;

        }
    };

    const submitName = async (name: string) => {
        const res = await setupEndpoints.updateName(name);
        navigateByStatus(res.status);
    };

    const submitImage = async (url: string) => {
        const res = await setupEndpoints.updateImage(url);
        navigateByStatus(res.status);
    };

    return {
        submitName,
        submitImage,
        navigateByStatus,
    };
};