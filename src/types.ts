export type AppRoute = {
  path: string;
  name: string;
  component: React.ComponentType<any>;
  isPrivate?: boolean;
};

export type AppParams = {
  [key: string]: string;
};

export type AppRouteProps = {
  route: AppRoute;
  params?: AppParams;
  navigation: any;
};

export type AppDispatch = (
  action: AppAction,
  dispatch: (action: AppAction) => void
) => void;