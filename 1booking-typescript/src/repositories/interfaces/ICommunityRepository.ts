import type { Community, PaginatedResult } from "@/types/index.js";

export interface ICommunityRepository {
  findByName(name: string):                  Promise<Community | null>;
  findWithMemberCount():                     Promise<Community[]>;
  findPosts(communityId: string):            Promise<unknown[]>;
  addMember(communityId: string, userId: string): Promise<void>;
  removeMember(communityId: string, userId: string): Promise<void>;
  search(query: CommunitySearchInput):       Promise<PaginatedResult<Community>>;
}

export interface CommunitySearchInput {
  category?:  string;
  isPrivate?: boolean;
  country?:   string;
  search?:    string;
  page:       number;
  limit:      number;
}
