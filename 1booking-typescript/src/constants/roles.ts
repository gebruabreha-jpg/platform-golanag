export enum Role {
  DIASPORA = "DIASPORA",
  LOCAL    = "LOCAL",
  MERCHANT = "MERCHANT",
  ADMIN    = "ADMIN",
}

export enum Permission {
  // Housekeeping
  READ_OWN_PROFILE     = "read_own_profile",
  UPDATE_OWN_PROFILE   = "update_own_profile",
  // Users
  READ_USERS           = "read_users",
  UPDATE_USERS         = "update_users",
  DELETE_USERS         = "delete_users",
  // Communities
  CREATE_COMMUNITY     = "create_community",
  MANAGE_ANY_COMMUNITY = "manage_any_community",
  // Housing
  CREATE_LISTING       = "create_listing",
  MANAGE_ANY_LISTING   = "manage_any_listing",
  // Marketplace
  SELL_ITEM            = "sell_item",
  BUY_ITEM             = "buy_item",
  MANAGE_ESCROW        = "manage_escrow",
  // Services
  LIST_SERVICE         = "list_service",
  BOOK_SERVICE         = "book_service",
  // Admin
  ADMIN_PANEL          = "admin_panel",
  VIEW_ANALYTICS       = "view_analytics",
}

/** Explicit permission map per role */
export const ROLE_PERMISSIONS: Record<Role, Permission[]> = {
  [Role.DIASPORA]: [
    Permission.READ_OWN_PROFILE, Permission.UPDATE_OWN_PROFILE,
    Permission.CREATE_COMMUNITY, Permission.BUY_ITEM, Permission.BOOK_SERVICE,
  ],
  [Role.LOCAL]: [
    Permission.READ_OWN_PROFILE, Permission.UPDATE_OWN_PROFILE,
    Permission.CREATE_COMMUNITY, Permission.BUY_ITEM, Permission.SELL_ITEM,
    Permission.CREATE_LISTING, Permission.BOOK_SERVICE,
  ],
  [Role.MERCHANT]: [
    Permission.READ_OWN_PROFILE, Permission.UPDATE_OWN_PROFILE,
    Permission.SELL_ITEM, Permission.CREATE_LISTING,
    Permission.BUY_ITEM, Permission.BOOK_SERVICE,
  ],
  [Role.ADMIN]: Object.values(Permission),
};

/** Role hierarchy – higher role inherits lower role permissions. */
export const ROLE_HIERARCHY: Role[] = [
  Role.LOCAL, Role.DIASPORA, Role.MERCHANT, Role.ADMIN,
];

/** Permissions for a given role or wildcard. */
export function getPermissionsForRole(role: Role | string): Permission[] {
  if (role === Role.ADMIN) return Object.values(Permission);
  const specific = ROLE_PERMISSIONS[role as Role];
  if (specific) return specific;
  return [];
}

export function hasPermission(role: string, permission: Permission): boolean {
  return getPermissionsForRole(role).includes(permission);
}

export function hasRole(role: string, ...roles: string[]): boolean {
  return roles.includes(role) || roles.includes("*");
}

export { Role as UserRole };