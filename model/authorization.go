// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

const (
	PERMISSION_LEVEL_SYSTEM  = "system"
	PERMISSION_LEVEL_TEAM    = "team"
	PERMISSION_LEVEL_CHANNEL = "channel"
)

type Permission struct {
	Id          string `json:"id"`
	Level       string `json:"level"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

var PERMISSION_INVITE_USER *Permission
var PERMISSION_ADD_USER_TO_TEAM *Permission
var PERMISSION_USE_SLASH_COMMANDS *Permission
var PERMISSION_MANAGE_SLASH_COMMANDS *Permission
var PERMISSION_MANAGE_OTHERS_SLASH_COMMANDS *Permission
var PERMISSION_CREATE_PUBLIC_CHANNEL *Permission
var PERMISSION_CREATE_PRIVATE_CHANNEL *Permission
var PERMISSION_MANAGE_PUBLIC_CHANNEL_MEMBERS *Permission
var PERMISSION_MANAGE_PRIVATE_CHANNEL_MEMBERS *Permission
var PERMISSION_ASSIGN_SYSTEM_ADMIN_ROLE *Permission
var PERMISSION_MANAGE_ROLES *Permission
var PERMISSION_MANAGE_TEAM_ROLES *Permission
var PERMISSION_MANAGE_CHANNEL_ROLES *Permission
var PERMISSION_CREATE_DIRECT_CHANNEL *Permission
var PERMISSION_MANAGE_PUBLIC_CHANNEL_PROPERTIES *Permission
var PERMISSION_MANAGE_PRIVATE_CHANNEL_PROPERTIES *Permission
var PERMISSION_LIST_TEAM_CHANNELS *Permission
var PERMISSION_JOIN_PUBLIC_CHANNELS *Permission
var PERMISSION_DELETE_PUBLIC_CHANNEL *Permission
var PERMISSION_DELETE_PRIVATE_CHANNEL *Permission
var PERMISSION_EDIT_OTHER_USERS *Permission
var PERMISSION_READ_CHANNEL *Permission
var PERMISSION_PERMANENT_DELETE_USER *Permission
var PERMISSION_UPLOAD_FILE *Permission
var PERMISSION_GET_PUBLIC_LINK *Permission
var PERMISSION_MANAGE_WEBHOOKS *Permission
var PERMISSION_MANAGE_OTHERS_WEBHOOKS *Permission
var PERMISSION_MANAGE_OAUTH *Permission
var PERMISSION_MANAGE_SYSTEM_WIDE_OAUTH *Permission
var PERMISSION_CREATE_POST *Permission
var PERMISSION_EDIT_POST *Permission
var PERMISSION_EDIT_OTHERS_POSTS *Permission
var PERMISSION_REMOVE_USER_FROM_TEAM *Permission
var PERMISSION_MANAGE_TEAM *Permission
var PERMISSION_IMPORT_TEAM *Permission

// General permission that encompases all system admin functions
// in the future this could be broken up to allow access to some
// admin functions but not others
var PERMISSION_MANAGE_SYSTEM *Permission

var ROLE_SYSTEM_USER *Role
var ROLE_SYSTEM_ADMIN *Role

var ROLE_TEAM_USER *Role
var ROLE_TEAM_ADMIN *Role

var ROLE_CHANNEL_USER *Role
var ROLE_CHANNEL_ADMIN *Role
var ROLE_CHANNEL_GUEST *Role

var BuiltInRoles map[string]*Role

func InitalizePermissions() {
	PERMISSION_INVITE_USER = &Permission{
		"invite_user",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.team_invite_user.name",
		"authentication.permissions.team_invite_user.description",
	}
	PERMISSION_ADD_USER_TO_TEAM = &Permission{
		"add_user_to_team",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.add_user_to_team.name",
		"authentication.permissions.add_user_to_team.description",
	}
	PERMISSION_USE_SLASH_COMMANDS = &Permission{
		"use_slash_commands",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.team_use_slash_commands.name",
		"authentication.permissions.team_use_slash_commands.description",
	}
	PERMISSION_MANAGE_SLASH_COMMANDS = &Permission{
		"manage_slash_commands",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_slash_commands.name",
		"authentication.permissions.manage_slash_commands.description",
	}
	PERMISSION_MANAGE_OTHERS_SLASH_COMMANDS = &Permission{
		"manage_others_slash_commands",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_others_slash_commands.name",
		"authentication.permissions.manage_others_slash_commands.description",
	}
	PERMISSION_CREATE_PUBLIC_CHANNEL = &Permission{
		"create_public_channel",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.create_public_channel.name",
		"authentication.permissions.create_public_channel.description",
	}
	PERMISSION_CREATE_PRIVATE_CHANNEL = &Permission{
		"create_private_channel",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.create_private_channel.name",
		"authentication.permissions.create_private_channel.description",
	}
	PERMISSION_MANAGE_PUBLIC_CHANNEL_MEMBERS = &Permission{
		"manage_public_channel_members",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.manage_public_channel_members.name",
		"authentication.permissions.manage_public_channel_members.description",
	}
	PERMISSION_MANAGE_PRIVATE_CHANNEL_MEMBERS = &Permission{
		"manage_private_channel_members",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.manage_private_channel_members.name",
		"authentication.permissions.manage_private_channel_members.description",
	}
	PERMISSION_ASSIGN_SYSTEM_ADMIN_ROLE = &Permission{
		"assign_system_admin_role",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.assign_system_admin_role.name",
		"authentication.permissions.assign_system_admin_role.description",
	}
	PERMISSION_MANAGE_ROLES = &Permission{
		"manage_roles",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.manage_roles.name",
		"authentication.permissions.manage_roles.description",
	}
	PERMISSION_MANAGE_TEAM_ROLES = &Permission{
		"manage_team_roles",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_team_roles.name",
		"authentication.permissions.manage_team_roles.description",
	}
	PERMISSION_MANAGE_CHANNEL_ROLES = &Permission{
		"manage_channel_roles",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.manage_channel_roles.name",
		"authentication.permissions.manage_channel_roles.description",
	}
	PERMISSION_MANAGE_SYSTEM = &Permission{
		"manage_system",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.manage_system.name",
		"authentication.permissions.manage_system.description",
	}
	PERMISSION_CREATE_DIRECT_CHANNEL = &Permission{
		"create_direct_channel",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.create_direct_channel.name",
		"authentication.permissions.create_direct_channel.description",
	}
	PERMISSION_MANAGE_PUBLIC_CHANNEL_PROPERTIES = &Permission{
		"manage__publicchannel_properties",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.manage_public_channel_properties.name",
		"authentication.permissions.manage_public_channel_properties.description",
	}
	PERMISSION_MANAGE_PRIVATE_CHANNEL_PROPERTIES = &Permission{
		"manage_private_channel_properties",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.manage_private_channel_properties.name",
		"authentication.permissions.manage_private_channel_properties.description",
	}
	PERMISSION_LIST_TEAM_CHANNELS = &Permission{
		"list_team_channels",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.list_team_channels.name",
		"authentication.permissions.list_team_channels.description",
	}
	PERMISSION_JOIN_PUBLIC_CHANNELS = &Permission{
		"join_public_channels",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.join_public_channels.name",
		"authentication.permissions.join_public_channels.description",
	}
	PERMISSION_DELETE_PUBLIC_CHANNEL = &Permission{
		"delete_public_channel",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.delete_public_channel.name",
		"authentication.permissions.delete_public_channel.description",
	}
	PERMISSION_DELETE_PRIVATE_CHANNEL = &Permission{
		"delete_private_channel",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.delete_private_channel.name",
		"authentication.permissions.delete_private_channel.description",
	}
	PERMISSION_EDIT_OTHER_USERS = &Permission{
		"edit_other_users",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.edit_other_users.name",
		"authentication.permissions.edit_other_users.description",
	}
	PERMISSION_READ_CHANNEL = &Permission{
		"read_channel",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.read_channel.name",
		"authentication.permissions.read_channel.description",
	}
	PERMISSION_PERMANENT_DELETE_USER = &Permission{
		"permanent_delete_user",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.permanent_delete_user.name",
		"authentication.permissions.permanent_delete_user.description",
	}
	PERMISSION_UPLOAD_FILE = &Permission{
		"upload_file",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.upload_file.name",
		"authentication.permissions.upload_file.description",
	}
	PERMISSION_GET_PUBLIC_LINK = &Permission{
		"get_public_link",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.get_public_link.name",
		"authentication.permissions.get_public_link.description",
	}
	PERMISSION_MANAGE_WEBHOOKS = &Permission{
		"manage_webhooks",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_webhooks.name",
		"authentication.permissions.manage_webhooks.description",
	}
	PERMISSION_MANAGE_OTHERS_WEBHOOKS = &Permission{
		"manage_others_webhooks",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_others_webhooks.name",
		"authentication.permissions.manage_others_webhooks.description",
	}
	PERMISSION_MANAGE_OAUTH = &Permission{
		"manage_oauth",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.manage_oauth.name",
		"authentication.permissions.manage_oauth.description",
	}
	PERMISSION_MANAGE_SYSTEM_WIDE_OAUTH = &Permission{
		"manage_sytem_wide_oauth",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.manage_sytem_wide_oauth.name",
		"authentication.permissions.manage_sytem_wide_oauth.description",
	}
	PERMISSION_CREATE_POST = &Permission{
		"create_post",
		PERMISSION_LEVEL_SYSTEM,
		"authentication.permissions.create_post.name",
		"authentication.permissions.create_post.description",
	}
	PERMISSION_EDIT_POST = &Permission{
		"edit_post",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.edit_post.name",
		"authentication.permissions.edit_post.description",
	}
	PERMISSION_EDIT_OTHERS_POSTS = &Permission{
		"edit_others_posts",
		PERMISSION_LEVEL_CHANNEL,
		"authentication.permissions.edit_others_posts.name",
		"authentication.permissions.edit_others_posts.description",
	}
	PERMISSION_REMOVE_USER_FROM_TEAM = &Permission{
		"remove_user_from_team",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.remove_user_from_team.name",
		"authentication.permissions.remove_user_from_team.description",
	}
	PERMISSION_MANAGE_TEAM = &Permission{
		"manage_team",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.manage_team.name",
		"authentication.permissions.manage_team.description",
	}
	PERMISSION_IMPORT_TEAM = &Permission{
		"import_team",
		PERMISSION_LEVEL_TEAM,
		"authentication.permissions.import_team.name",
		"authentication.permissions.import_team.description",
	}
}

func InitalizeRoles() {
	InitalizePermissions()
	BuiltInRoles = make(map[string]*Role)

	ROLE_CHANNEL_USER = &Role{
		"channel_user",
		"authentication.roles.channel_user.name",
		"authentication.roles.channel_user.description",
		[]string{
			PERMISSION_READ_CHANNEL.Id,
			PERMISSION_MANAGE_PUBLIC_CHANNEL_MEMBERS.Id,
			PERMISSION_MANAGE_PRIVATE_CHANNEL_MEMBERS.Id,
			PERMISSION_UPLOAD_FILE.Id,
			PERMISSION_GET_PUBLIC_LINK.Id,
			PERMISSION_CREATE_POST.Id,
			PERMISSION_EDIT_POST.Id,
			PERMISSION_USE_SLASH_COMMANDS.Id,
		},
	}
	BuiltInRoles[ROLE_CHANNEL_USER.Id] = ROLE_CHANNEL_USER
	ROLE_CHANNEL_ADMIN = &Role{
		"channel_admin",
		"authentication.roles.channel_admin.name",
		"authentication.roles.channel_admin.description",
		[]string{
			PERMISSION_MANAGE_CHANNEL_ROLES.Id,
		},
	}
	BuiltInRoles[ROLE_CHANNEL_ADMIN.Id] = ROLE_CHANNEL_ADMIN
	ROLE_CHANNEL_GUEST = &Role{
		"guest",
		"authentication.roles.global_guest.name",
		"authentication.roles.global_guest.description",
		[]string{},
	}
	BuiltInRoles[ROLE_CHANNEL_GUEST.Id] = ROLE_CHANNEL_GUEST

	ROLE_TEAM_USER = &Role{
		"team_user",
		"authentication.roles.team_user.name",
		"authentication.roles.team_user.description",
		[]string{
			PERMISSION_LIST_TEAM_CHANNELS.Id,
			PERMISSION_JOIN_PUBLIC_CHANNELS.Id,
		},
	}
	BuiltInRoles[ROLE_TEAM_USER.Id] = ROLE_TEAM_USER
	ROLE_TEAM_ADMIN = &Role{
		"team_admin",
		"authentication.roles.team_admin.name",
		"authentication.roles.team_admin.description",
		[]string{
			PERMISSION_EDIT_OTHERS_POSTS.Id,
			PERMISSION_ADD_USER_TO_TEAM.Id,
			PERMISSION_REMOVE_USER_FROM_TEAM.Id,
			PERMISSION_MANAGE_TEAM.Id,
			PERMISSION_IMPORT_TEAM.Id,
			PERMISSION_MANAGE_TEAM_ROLES.Id,
			PERMISSION_MANAGE_CHANNEL_ROLES.Id,
			PERMISSION_MANAGE_OTHERS_WEBHOOKS.Id,
			PERMISSION_MANAGE_SLASH_COMMANDS.Id,
			PERMISSION_MANAGE_OTHERS_SLASH_COMMANDS.Id,
			PERMISSION_MANAGE_WEBHOOKS.Id,
		},
	}
	BuiltInRoles[ROLE_TEAM_ADMIN.Id] = ROLE_TEAM_ADMIN

	ROLE_SYSTEM_USER = &Role{
		"system_user",
		"authentication.roles.global_user.name",
		"authentication.roles.global_user.description",
		[]string{
			PERMISSION_CREATE_DIRECT_CHANNEL.Id,
			PERMISSION_PERMANENT_DELETE_USER.Id,
			PERMISSION_MANAGE_OAUTH.Id,
		},
	}
	BuiltInRoles[ROLE_SYSTEM_USER.Id] = ROLE_SYSTEM_USER
	ROLE_SYSTEM_ADMIN = &Role{
		"system_admin",
		"authentication.roles.global_admin.name",
		"authentication.roles.global_admin.description",
		// System admins can do anything channel and team admins can do
		// plus everything members of teams and channels can do to all teams
		// and channels on the system
		append(
			append(
				append(
					append(
						[]string{
							PERMISSION_ASSIGN_SYSTEM_ADMIN_ROLE.Id,
							PERMISSION_MANAGE_ROLES.Id,
							PERMISSION_MANAGE_SYSTEM.Id,
							PERMISSION_MANAGE_PUBLIC_CHANNEL_PROPERTIES.Id,
							PERMISSION_DELETE_PUBLIC_CHANNEL.Id,
							PERMISSION_CREATE_PUBLIC_CHANNEL.Id,
							PERMISSION_MANAGE_PRIVATE_CHANNEL_PROPERTIES.Id,
							PERMISSION_DELETE_PRIVATE_CHANNEL.Id,
							PERMISSION_CREATE_PRIVATE_CHANNEL.Id,
							PERMISSION_MANAGE_SYSTEM_WIDE_OAUTH.Id,
							PERMISSION_MANAGE_OTHERS_WEBHOOKS.Id,
							PERMISSION_EDIT_OTHER_USERS.Id,
							PERMISSION_MANAGE_OAUTH.Id,
							PERMISSION_INVITE_USER.Id,
						},
						ROLE_TEAM_USER.Permissions...,
					),
					ROLE_CHANNEL_USER.Permissions...,
				),
				ROLE_TEAM_ADMIN.Permissions...,
			),
			ROLE_CHANNEL_ADMIN.Permissions...,
		),
	}
	BuiltInRoles[ROLE_SYSTEM_ADMIN.Id] = ROLE_SYSTEM_ADMIN

}

func RoleIdsToString(roles []string) string {
	output := ""
	for _, role := range roles {
		output += role + ", "
	}

	if output == "" {
		return "[<NO ROLES>]"
	}

	return output[:len(output)-1]
}

func init() {
	InitalizeRoles()
}
